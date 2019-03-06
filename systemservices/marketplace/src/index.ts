import { service as MESG } from "mesg-js"
import { TaskInputs, Service, EmitEventReply } from "mesg-js/lib/service"

import Web3 from "web3"
import { EventLog } from "web3/types"

import { newBlockEventEmitter } from "./newBlock"

import marketplaceABI from "./contracts/Marketplace.abi.json"
import { Marketplace } from "./contracts/Marketplace"

import ERC20ABI from "./contracts/ERC20.abi.json"
import { ERC20 } from "./contracts/ERC20"

import serviceCreated from "./events/serviceCreated"
import serviceOfferCreated from "./events/serviceOfferCreated"
import serviceOfferDisabled from "./events/serviceOfferDisabled"
import serviceOwnershipTransferred from "./events/serviceOwnershipTransferred"
import serviceVersionCreated from "./events/serviceVersionCreated"
import servicePurchased from "./events/servicePurchased"

import publishServiceVersion from "./tasks/publishServiceVersion"
import getService from "./tasks/getService"
import createServiceOffer from "./tasks/createServiceOffer"
import disableServiceOffer from "./tasks/disableServiceOffer"
import listServices from "./tasks/listServices"
import purchase from "./tasks/purchase"
import sendSignedTransaction from "./tasks/sendSignedTransaction"
import transferServiceOwnership from "./tasks/transferServiceOwnership"
import checkForDeployment from "./tasks/checkForDeployment"
import Contract from "web3/eth/contract";
import { createTransactionTemplate } from "./contracts/utils";

const providerEndpoint = process.env.PROVIDER_ENDPOINT as string
const marketplaceAddress = process.env.MARKETPLACE_ADDRESS
const ERC20Address = process.env.TOKEN_ADDRESS
const blockConfirmations = parseInt(<string>process.env.BLOCK_CONFIRMATIONS, 10)
const defaultGas = parseInt(<string>process.env.DEFAULT_GAS, 10)
const pollingTime = parseInt(<string>process.env.POLLING_TIME, 10)

const eventHandlers: {[key: string]: (mesg: Service, event: EventLog) => Promise<EmitEventReply | Error>} = {
  'ServiceCreated': serviceCreated,
  'ServiceOfferCreated': serviceOfferCreated,
  'ServiceOfferDisabled': serviceOfferDisabled,
  'ServiceOwnershipTransferred': serviceOwnershipTransferred,
  'ServicePurchased': servicePurchased,
  'ServiceVersionCreated': serviceVersionCreated,
}

const main = async () => {
  const mesg = MESG()
  const web3 = new Web3(providerEndpoint)
  const marketplace = new web3.eth.Contract(marketplaceABI, marketplaceAddress) as Marketplace
  const token = new web3.eth.Contract(ERC20ABI, ERC20Address) as ERC20

  const chainID = await web3.eth.net.getId()
  console.log('chainID', chainID)
  const defaultGasPrice = await web3.eth.getGasPrice()
  console.log('defaultGasPrice', defaultGasPrice)

  const createTransaction = createTransactionTemplate(chainID, web3, defaultGas, defaultGasPrice)

  mesg.listenTask({
    listServices: listServices(marketplace),
    getService: getService(marketplace),
    publishServiceVersion: publishServiceVersion(marketplace, createTransaction),
    createServiceOffer: createServiceOffer(marketplace, createTransaction),
    disableServiceOffer: disableServiceOffer(marketplace, createTransaction),
    purchase: purchase(marketplace, token, createTransaction),
    transferServiceOwnership: transferServiceOwnership(marketplace, createTransaction),
    sendSignedTransaction: sendSignedTransaction(web3),
    checkForDeployment: checkForDeployment(marketplace),
  })
  .on('error', error => console.error('catch listenTask', error))

  const newBlock = await newBlockEventEmitter(web3, blockConfirmations, null, pollingTime)
  newBlock.on('newBlock', async blockNumber => {
    try {
      console.error('new block', blockNumber)
    
      const events = await marketplace.getPastEvents("allEvents", {
        fromBlock: blockNumber,
        toBlock: blockNumber,
      })
      events.forEach(async event => {
        // TODO: check if really async
        try {
          if (!eventHandlers[event.event]) {
            return console.error('Event not implemented', event.event)
          }
          eventHandlers[event.event](mesg, event)
        } catch(error) {
          return console.error('An error occurred during processing of an event', event)
        }
      })
    }
    catch (error) {
      console.error('catch newBlock on', error)
    }
  })

  console.log('service is ready and running')
}

try {
  main()
    .catch(error => console.error('catch promise', error))
} catch (error) {
  console.error('catch try', error)
}
