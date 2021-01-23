import { useContext, useEffect, useState } from "react";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import { makeStyles } from "@material-ui/core/styles";
import { Typography, Card, CardContent } from "@material-ui/core";
import {
  AreaChart,
  Area,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import { getNetworkMetrics } from "../api/NetworkMetricsApi";
import dayjs from "dayjs";
import { formatBytes } from "../utils/utils";

const useStyles = makeStyles({
  root: {
    marginBottom: 24,
  },
});

const MetricsChart = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  const classes = useStyles();

  const [networkMetricsData, setNetworkMetricsData] = useState([]);

  const updateNetworkMetrics = () => {
    getNetworkMetrics()
      .then(function (response) {
        let newMetricsData = [];

        appStore.networkMetrics.data = response?.data;

        for (let metric of response?.data) {
          let parsedMetrics = {};

          parsedMetrics["timestamp"] = dayjs(metric?.timestamp).format(
            "HH:mm:ss A"
          );
          parsedMetrics["upload"] = parseFloat(
            metric?.network_metrics?.groupTraffic?.transmitSpeedBps
          );
          parsedMetrics["download"] = parseFloat(
            metric?.network_metrics?.groupTraffic?.receiveSpeedBps
          );

          newMetricsData.push(parsedMetrics);
        }
        setNetworkMetricsData(newMetricsData.reverse());
      })
      .catch(function (err) {
        console.error(err);
      });
  };

  useEffect(() => {
    setInterval(updateNetworkMetrics, 1000);
  }, []);

  const ToolTipFormatter = (value, name, props) => {
    return formatBytes(value);
  };

  return (
    <div>
      <Card className={classes.root}>
        <CardContent>
          <Typography variant="h6">Network Usage</Typography>
          <div style={{ height: 400 }}>
            <ResponsiveContainer>
              <AreaChart
                data={networkMetricsData}
                margin={{ top: 30, right: 0, left: 0, bottom: 50 }}
              >
                <defs>
                  <linearGradient
                    id="colorDownload"
                    x1="0"
                    y1="0"
                    x2="0"
                    y2="1"
                  >
                    <stop offset="5%" stopColor="#82ca9d" stopOpacity={0.8} />
                    <stop offset="95%" stopColor="#82ca9d" stopOpacity={0} />
                  </linearGradient>
                  <linearGradient id="colorUpload" x1="0" y1="0" x2="0" y2="1">
                    <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8} />
                    <stop offset="95%" stopColor="#8884d8" stopOpacity={0} />
                  </linearGradient>
                </defs>
                <XAxis
                  dataKey="timestamp"
                  interval={0}
                  angle={-45}
                  textAnchor="end"
                />
                <YAxis />
                <Tooltip formatter={ToolTipFormatter} />
                <Area
                  type="monotone"
                  dataKey="download"
                  stroke="#82ca9d"
                  fillOpacity={1}
                  fill="url(#colorDownload)"
                />
                <Area
                  type="monotone"
                  dataKey="upload"
                  stroke="#8884d8"
                  fillOpacity={1}
                  fill="url(#colorUpload)"
                />
              </AreaChart>
            </ResponsiveContainer>
          </div>
        </CardContent>
      </Card>
    </div>
  );
});

export default MetricsChart;
