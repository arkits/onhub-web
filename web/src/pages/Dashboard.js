import { useContext, useEffect, useState } from "react";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import StatusBanner from "../components/StatusBanner";
import DeviceCards from "../components/DeviceCards";

const Dashboard = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  if (appStore.isInitialLoadComplete) {
    return (
      <>
        <StatusBanner />
        <DeviceCards />
      </>
    );
  } else {
    return null;
  }
});

export default Dashboard;
