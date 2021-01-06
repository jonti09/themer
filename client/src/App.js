import { useEffect, useState } from "react";
import { HomeScreen } from "./screens/HomeScreen";
import { callAPI } from "./utils";

function App() {
  const [themes, setThemes] = useState([]);

  useEffect(() => {
    const getThemes = async () => {
      const { themes } = await callAPI({ url: "/themes/", method: "GET" });
      setThemes(themes);
    };

    getThemes();
  }, []);

  return <HomeScreen themes={themes} />;
}

export default App;
