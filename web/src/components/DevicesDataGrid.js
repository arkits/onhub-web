import React, { useContext } from "react";
import { DataGrid } from "@material-ui/data-grid";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import { Paper } from "@material-ui/core";

const columns = [
  { field: "id", headerName: "ID", width: 80 },
  { field: "name", headerName: "Device Name", width: 290 },
  { field: "connectionType", headerName: "Type", width: 130 },
  { field: "ipAddress", headerName: "IP", width: 130 },
  { field: "upload", headerName: "UP", width: 130 },
  { field: "download", headerName: "DOWN", width: 130 },
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
                  field: "download",
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
