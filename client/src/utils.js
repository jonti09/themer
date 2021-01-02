import axios from "axios";

const getBaseURL = () => {
  let baseUrl;
  if (process.env.NODE_ENV === "development") {
    baseUrl =
      window.location.protocol +
      "//" +
      window.location.hostname +
      `:8000${process.env.VUE_APP_SERVER_SUFFIX || ""}`;
  } else {
    if (!window.location.origin) {
      window.location.origin =
        window.location.protocol +
        "//" +
        window.location.hostname +
        (window.location.port ? ":" + window.location.port : "");
    }
    baseUrl =
      window.location.origin + `${process.env.VUE_APP_SERVER_SUFFIX || ""}`;
  }
  return baseUrl;
};

export const callAPI = async ({ url, method }) => {
  const { data } = await axios({
    baseURL: getBaseURL(),
    url,
    method,
  });
  return data;
};
