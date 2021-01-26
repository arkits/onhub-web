import axios from "axios";
import { API_BASE } from "./ApiConstants";

function getDebugVersion() {
  return axios({
    method: "GET",
    url: API_BASE,
  });
}

export { getDebugVersion };
