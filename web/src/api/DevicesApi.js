import axios from "axios";
import { API_BASE } from "./ApiConstants";

function getDevices() {
  return axios({
    method: "GET",
    url: API_BASE + "/devices",
  });
}

export { getDevices };
