import Service, { EmitEventReply } from "mesg-js/lib/service/service"
import { EventLog } from "web3/types";
import { hexToAscii } from "../contracts/utils";

export = (mesg: Service, event: EventLog): Promise<EmitEventReply | Error> => {
  return mesg.emitEvent('serviceCreated', {
    sid: hexToAscii(event.returnValues.sid),
    hashedSid: event.returnValues.hashedSid,
    owner: event.returnValues.owner,
  })
}
