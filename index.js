const axios = require("axios");
const node_uid = require("node-uid");

(async function run() {
  for (let index = 0; index < 1; index++) {
    const id = node_uid(10);
    const p = axios.post("http://localhost:8090/charges/", {
      order_id: "770434815" + id
    });

    const q = axios.post("http://localhost:8090/charges/", {
      order_id: "770434815" + id
    });

    try {
      await Promise.all([p, q]);
    } catch (error) {
      console.log(error);
    }
  }
})();
