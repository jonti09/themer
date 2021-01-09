import { Body } from "./Body";
import style from "./Card.module.scss";
import { Header } from "./Header";

export const Card = ({ theme, isActive, applyTheme }) => {
  return (
    <div className={style.card}>
      <Header name={theme.name} label={theme.label} applyTheme={applyTheme} />
      <Body theme={theme} isActive={isActive} />
    </div>
  );
};
