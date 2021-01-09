import { useEffect, useState } from "react";
import { HomeScreen } from "./screens/HomeScreen";
import { callAPI } from "./utils";

function App() {
  const [themes, setThemes] = useState([]);

  useEffect(() => {
    const getThemes = async () => {
      try {
        const { themes } = await callAPI({ url: "/themes/", method: "GET" });
        setThemes(themes);
      } catch (e) {}
    };

    getThemes();
  }, []);

  const applyTheme = async (themeName) => {
    await callAPI({
      url: "/apply/",
      method: "POST",
      formData: { theme: themeName },
    });
  };

  return <HomeScreen themes={themes} applyTheme={applyTheme} />;
}

export default App;
