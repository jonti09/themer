import { Card } from "../components/terminal-card";
import style from "./HomeScreen.module.scss";

export const HomeScreen = ({ themes }) => {
  const card = (theme) => {
    return <Card key={theme.name} theme={theme} className={style.card} />;
  };
  return (
    <div className={style.container}>{themes.map((theme) => card(theme))}</div>
  );
};
