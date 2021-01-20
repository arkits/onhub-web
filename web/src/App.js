import { useEffect, useState } from "react";
import Container from "@material-ui/core/Container";
import axios from "axios";

function DeviceCards({ devices }) {
  let elements = [];

  devices.forEach((element) => {
    elements.push(<h4 key={element.id}>{element.friendlyName}</h4>);
  });

  return elements;
}

function App() {
  const [devices, setDevices] = useState([]);

  useEffect(() => {
    // Update the document title using the browser API
    axios({
      method: "get",
      url: "http://localhost:4209/ohw/api/devices",
      headers: {},
    })
      .then(function (response) {
        setDevices(response?.data);
      })
      .catch(function (err) {
        console.error(err);
      });
  }, []);

  return (
    <div className="App">
      <Container maxWidth="md">
        <h1>OnHub-Web</h1>
        <h2>Number of Connected Devices - {devices.length}</h2>
        <DeviceCards devices={devices} />
      </Container>
    </div>
  );
}

export default App;
