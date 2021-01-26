import axios from "axios";
import { API_BASE } from "./ApiConstants";

function getVersion() {
  return axios({
    method: "GET",
    url: API_BASE,
  });
}

export { getVersion };
