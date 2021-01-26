import { useContext, useEffect, useState } from "react";
import {
  Typography,
  LinearProgress,
  Card,
  CardContent,
  Container,
} from "@material-ui/core";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import { getDebugVersion } from "../api/DebugApi";
import { getDevices } from "../api/DevicesApi";
import { getNetworkMetrics } from "../api/NetworkMetricsApi";
import dayjs from "dayjs";

function ErrorBanner({ apiError }) {
  if (apiError != null) {
    return (
      <>
        <Card style={{ margin: 20, background: "red" }}>
          <CardContent>
            <Typography variant="h5">
              Server Error: {apiError.message}
            </Typography>
            <Typography variant="body1">
              An Error with the Server has occured. Here are some details -
            </Typography>
            <pre>{JSON.stringify(apiError, null, 4)}</pre>
          </CardContent>
        </Card>
      </>
    );
  } else {
    return null;
  }
}

const InitialDataLoader = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  const [isLoading, setIsLoading] = useState(false);

  const [apiError, setApiError] = useState(null);

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
        appStore.networkMetrics.parsedData = newMetricsData.reverse();

        setIsLoading(false);
        appStore.isInitialLoadComplete = true;
      })
      .catch(function (err) {
        console.error(err);
      });
  };

  useEffect(() => {
    setIsLoading(true);
    async function fetchInitialData() {
      try {
        let debugVersion = await getDebugVersion();
        appStore.appVersion = debugVersion?.data;

        let devices = await getDevices();
        appStore.devices = devices?.data;

        setInterval(() => {
          updateNetworkMetrics();
        }, 1000);
      } catch (error) {
        console.error(error);
        setIsLoading(false);
        setApiError(error);
      }
    }
    fetchInitialData();
  }, []);

  if (isLoading) {
    return (
      <Container maxWidth="md">
        <div style={{ marginTop: "10%" }}>
          <center>
            <Typography variant="h5">Loading Initial Data...</Typography> <br />
            <LinearProgress color="secondary" />
          </center>
        </div>
      </Container>
    );
  } else {
    return <ErrorBanner apiError={apiError} />;
  }
});

export default InitialDataLoader;
