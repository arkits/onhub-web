import { Fragment, useContext } from "react";
import { observer } from "mobx-react";
import { AppStoreContext } from "../store/AppStore";
import { Container } from "@material-ui/core";
import StatusBanner from "../components/StatusBanner";
import DevicesDataGrid from "../components/DevicesDataGrid";
import MetricsChart from "../components/MetricsChart";
import Footer from "../components/Footer";

const Dashboard = observer(({}) => {
  const appStore = useContext(AppStoreContext);

  if (appStore.isInitialLoadComplete) {
    return (
      <Fragment>
        <Container maxWidth="md">
          <div style={{ marginTop: "10%", paddingBottom: "20px" }}>
            <StatusBanner />
          </div>
        </Container>

        <Container maxWidth="lg">
          <MetricsChart />
        </Container>

        <Container maxWidth="md">
          <DevicesDataGrid />
        </Container>
        <Footer />
      </Fragment>
    );
  } else {
    return null;
  }
});

export default Dashboard;
