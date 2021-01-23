import React from "react";
import { Card, CardContent } from "@material-ui/core";

function ApiResponseCard({ apiResponse }) {
  if (apiResponse.show) {
    let background = "#aa2e25";
    if (apiResponse.type === "success") {
      background = "#357a38";
    }
    return (
      <React.Fragment>
        <Card style={{ background: background }}>
          <CardContent>{apiResponse.message}</CardContent>
        </Card>
      </React.Fragment>
    );
  } else {
    return null;
  }
}

export default ApiResponseCard;
