import axios from "axios";

const stuckoverflowAPI = axios.create({
  baseURL: "https://stuckoverflow.herokuapp.com",
});

export default stuckoverflowAPI;