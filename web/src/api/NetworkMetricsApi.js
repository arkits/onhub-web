import axios from "axios";
import { API_BASE } from "./ApiConstants";

function getNetworkMetricsStatus() {
  return axios({
    method: "GET",
    url: API_BASE + "/network-metrics/status",
  });
}

function startPollingForNetworkMetrics() {
  return axios({
    method: "POST",
    url: API_BASE + "/network-metrics/start-polling",
  });
}

function getNetworkMetrics() {
  return axios({
    method: "GET",
    url: API_BASE + "/network-metrics",
  });
}

export {
  getNetworkMetricsStatus,
  startPollingForNetworkMetrics,
  getNetworkMetrics,
};
