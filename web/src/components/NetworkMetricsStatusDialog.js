import React, { useContext, useState, forwardRef, useEffect } from "react";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Slide,
} from "@material-ui/core";
import {
  getNetworkMetricsStatus,
  startPollingForNetworkMetrics,
} from "../api/NetworkMetricsApi";
import ApiResponseCard from "./ApiResponseCard";

const Transition = forwardRef(function Transition(props, ref) {
  return <Slide direction="up" ref={ref} {...props} />;
});

const NetworkMetricsStatusDialog = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  const networkMetricsStatus = appStore.networkMetrics?.status;

  const [open, setOpen] = useState(false);

  const [apiResponse, setApiResponse] = useState({
    show: false,
    type: "error",
    message: null,
  });

  const handleDialogOpen = () => {
    setOpen(true);
  };

  const handleDialogClose = () => {
    setOpen(false);
  };

  const beginNetworkMetricsCollection = () => {
    startPollingForNetworkMetrics()
      .then(function (response) {
        setApiResponse({
          show: true,
          type: "success",
          message: response?.data?.message,
        });
      })
      .catch(function (err) {
        setApiResponse({
          show: true,
          message: err,
        });
      });
  };

  const refreshNetworkMetricsStatus = () => {
    getNetworkMetricsStatus()
      .then(function (response) {
        appStore.networkMetrics.status = response?.data;
      })
      .catch(function (err) {
        setApiResponse({
          show: true,
          message: err,
        });
      });
  };

  useEffect(() => {
    getNetworkMetricsStatus()
      .then(function (response) {
        appStore.networkMetrics.status = response?.data;
      })
      .catch(function (err) {
        console.error(err);
      });
  }, []);

  return (
    <div>
      <Button color="inherit" onClick={handleDialogOpen}>
        Status
      </Button>
      <Dialog
        open={open}
        TransitionComponent={Transition}
        keepMounted
        onClose={handleDialogClose}
        aria-labelledby="network-metrics-status-dialog-title"
        aria-describedby="network-metrics-status-dialog-description"
      >
        <DialogTitle id="network-metrics-status-dialog-title">
          Network Metrics Status
        </DialogTitle>
        <DialogContent>
          <pre>{JSON.stringify(networkMetricsStatus, null, 4)}</pre>
          <ApiResponseCard apiResponse={apiResponse} />
        </DialogContent>
        <DialogActions>
          {!appStore?.networkMetrics?.status?.is_polling && (
            <Button onClick={beginNetworkMetricsCollection} color="primary">
              Begin Network Metrics Collection
            </Button>
          )}
          <Button onClick={refreshNetworkMetricsStatus} color="primary">
            Refresh Stats
          </Button>
          <Button onClick={handleDialogClose} color="primary">
            Close
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
});

export default NetworkMetricsStatusDialog;
