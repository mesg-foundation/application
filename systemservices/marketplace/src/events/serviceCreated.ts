import Service, { EmitEventReply } from "mesg-js/lib/service/service"
import { EventLog } from "web3/types";

export = (mesg: Service, event: EventLog): Promise<EmitEventReply | Error> => {
  return mesg.emitEvent('serviceCreated', {
    sid: event.returnValues.sid, // TODO: to convert to ascii
    owner: event.returnValues.owner,
  })
}
