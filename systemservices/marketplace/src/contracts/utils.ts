import Web3 from 'web3'
import BigNumber from 'bignumber.js';
import Contract from 'web3/eth/contract';
import { TaskInputs } from 'mesg-js/lib/service';

BigNumber.config({ EXPONENTIAL_AT: 100 })

const hexToAscii = (x: string) => {
  if (!x) return ''
  return Web3.utils.hexToAscii(x).replace(/\u0000/g, '')
}

const asciiToHex = (x: string) => Web3.utils.asciiToHex(x)
const sha3 = (x: string) => Web3.utils.sha3(x)

const toUnit = (x: string|BigNumber): string => {
  const n = new BigNumber(x).times(1e18)
  if (!n.integerValue().eq(n)) {
    throw new Error('Number of decimals of ' + x + ' is higher than 18')
  }
  return n.toString()
}
const fromUnit = (x: string|BigNumber) => new BigNumber(x).dividedBy(1e18)

const parseTimestamp = (x: string) => new Date(new BigNumber(x).times(1000).toNumber())

const hashToHex = (x: string) => {
  if (x.startsWith('0x')) {
    return x
  }
  return '0x' + x
}
const hexToHash = (x: string) => x.replace(/^0x/g, '')//.replace(/0*$/g, '')

interface CreateTransaction {
  (
    contract: Contract,
    inputs: TaskInputs,
    data: string,
    shiftNonce?: number
  ): Promise<any>;
};

const createTransactionTemplate = (
  chainID: number,
  web3: Web3,
  defaultGas: number,
  defaultGasPrice: number
): CreateTransaction => async (
  contract: Contract,
  inputs: TaskInputs,
  data: string,
  shiftNonce?: number
) => {
  return {
    chainID: chainID,
    to: contract.options.address,
    nonce: (await web3.eth.getTransactionCount(inputs.from)) + (shiftNonce || 0),
    gas: inputs.gas || defaultGas,
    gasPrice: inputs.gasPrice || defaultGasPrice,
    value: '0',
    data: data
  }
}

export {
  hexToAscii,
  asciiToHex,
  sha3,
  toUnit,
  fromUnit,
  parseTimestamp,
  createTransactionTemplate,
  CreateTransaction,
  hashToHex,
  hexToHash,
}
