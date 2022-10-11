import { googleRPC } from 'reactrpc';
import messages from "../proto/cryptoData_pb";
import services from "../proto/cryptoData_grpc_web_pb";


const URL = "http://" + window.location.hostname + "8080";
googleRPC.build(messages, services, URL);




export default googleRPC.wrapper(<streamingFunctions/>);