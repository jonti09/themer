import style from "./Header.module.scss";

export const Header = ({ name }) => {
  return (
    <div className={style.container}>
      <h5 className={style.header}>{name}</h5>
      <div className={style.buttons}>
        <div className={style.minimize}></div>
        <div className={style.maximize}></div>
        <div className={style.close}></div>
      </div>
    </div>
  );
};
