const express = require("express");
const _ = require("lodash");

const app = express();

app.get("/health", (req, res) => {
  res.json({ status: "ok", project: _.kebabCase("Project Two") });
});

app.listen(3000, () => console.log("project2 listening on 3000"));
