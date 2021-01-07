import { Body } from "./Body";
import style from "./Card.module.scss";
import { Header } from "./Header";

export const Card = ({ theme, isActive }) => {
  return (
    <div className={style.card}>
      <Header name={theme.name}></Header>
      <Body theme={theme} isActive={isActive} />
    </div>
  );
};
