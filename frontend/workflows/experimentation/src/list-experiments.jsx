import React, { useState } from "react";
import { client, Row, Table } from "@clutch-sh/core";
import { Button, Container } from "@material-ui/core";

import { StartAbortExperiment, StartLatencyExperiment } from "./start-experiment";

function isFunction(functionToCheck) {
  return functionToCheck && {}.toString.call(functionToCheck) === '[object Function]';
 }

function renderListItem(experiment, columns, mapping) {
  var data = [];
  var item;
  console.log(experiment)
  columns.forEach(column => {
    let value = mapping[column]
    if (isFunction(value)) {
      item = value(experiment);
    } else {
      item = experiment[value];
    }

    data.push(item);
  });

  return (
    <Row
      key={experiment.id}
      data={data}
    />
  );
}

const ListExperiments = ({ heading, columns, mapping }) => {
  console.log('heading: ' + heading);
  console.log(JSON.stringify(columns, null, 4));
  console.log(JSON.stringify(mapping, null, 4));
  console.log('RA: TEST');
  const [experiments, setExperiments] = useState();

  let column_names = columns.map(name => name.toUpperCase()  );

  if (!experiments) {
    client.post("/v1/experiments/get").then(response => {
      setExperiments(response?.data?.experiments || []);
    });

    return (
      <Table headings={column_names} />
    );
  }

  return (
    <Container>
      <Table
        data={experiments}
        headings={column_names}
      >
        {experiments.map(e => {
          return renderListItem(e, columns, mapping);
        })}
      </Table>
      <Button onClick={StartAbortExperiment}>Start Abort Experiment</Button>
      <Button onClick={StartLatencyExperiment}>Start Latency Experiment</Button>
    </Container>
  );
};

export default ListExperiments;
