import React, { useContext } from "react";
import { DataGrid } from "@material-ui/data-grid";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import { Paper } from "@material-ui/core";
import { formatBytes } from "../utils/utils";

const columns = [
  { field: "name", headerName: "Device Name", width: 280 },
  { field: "upload", headerName: "UP", width: 130, hide: true },
  {
    field: "prettyUpload",
    headerName: "UP",
    width: 130,
    valueGetter: (params) => formatBytes(params.getValue("upload")),
    sortComparator: (v1, v2, param1, param2) =>
      param1.row.upload - param2.row.upload,
  },
  { field: "download", headerName: "DOWN", width: 130, hide: true },
  {
    field: "prettyDownload",
    headerName: "DOWN",
    width: 130,
    valueGetter: (params) => formatBytes(params.getValue("download")),
    sortComparator: (v1, v2, param1, param2) =>
      param1.row.download - param2.row.download,
  },
  { field: "connectionType", headerName: "CONN", width: 130 },
  { field: "ipAddress", headerName: "IP", width: 130 },
];

const DevicesDataGrid = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  const networkMetricsData = appStore?.networkMetrics?.data;

  if (networkMetricsData) {
    const stationMetrics =
      networkMetricsData[0]?.network_metrics?.stationMetrics;

    let rows = [];

    for (const stationMetric of stationMetrics) {
      let row = {
        id: stationMetric?.station?.id,
        name: stationMetric?.station?.friendlyName,
        connectionType: stationMetric?.station?.connectionType,
        connected: stationMetric?.station?.connected,
        ipAddress: stationMetric?.station?.ipAddress,
        upload: parseFloat(stationMetric?.traffic?.transmitSpeedBps) || 0,
        download: parseFloat(stationMetric?.traffic?.receiveSpeedBps) || 0,
      };

      rows.push(row);
    }

    return (
      <>
        <Paper>
          <div style={{ height: 800, width: "100%" }}>
            <DataGrid
              sortModel={[
                {
                  field: "prettyDownload",
                  sort: "desc",
                },
              ]}
              rows={rows}
              columns={columns}
            />
          </div>
        </Paper>
        <br /> <br />
      </>
    );
  } else {
    return null;
  }
});

export default DevicesDataGrid;
