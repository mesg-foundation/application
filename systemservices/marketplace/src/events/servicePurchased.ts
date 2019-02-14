import Service, { EmitEventReply } from "mesg-js/lib/service/service"
import { EventLog } from "web3/types";

export = (mesg: Service, event: EventLog): Promise<EmitEventReply | Error> => {
  return mesg.emitEvent('servicePurchased', {
    sidHash: event.returnValues.sidHash,
    offerIndex: event.returnValues.offerIndex,
    purchaser: event.returnValues.purchaser,
    price: event.returnValues.price,
    duration: event.returnValues.duration,
    expire: event.returnValues.expire,
  })
}
