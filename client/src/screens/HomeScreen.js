import { Card } from "../components/terminal-card";

export const HomeScreen = ({ themes }) => {
  const card = (theme) => {
    return <Card key={theme.name} theme={theme} />;
  };
  return <>{themes.map((theme) => card(theme))}</>;
};
