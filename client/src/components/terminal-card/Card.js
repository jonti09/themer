import { Body } from "./Body";
import style from "./Card.module.scss";
import { Header } from "./Header";

export const Card = ({ theme }) => {
  return (
    <div className={style.card}>
      <Header name={theme.name}></Header>
      <Body theme={theme} />
    </div>
  );
};
