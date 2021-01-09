import { useState } from "react";
import { ReactComponent as ApplyIcon } from "../../assets/icons/apply.svg";
import { ReactComponent as CheckIcon } from "../../assets/icons/check.svg";
import style from "./Header.module.scss";

export const Header = ({ name, label, applyTheme }) => {
  const [applied, setApplied] = useState(false);

  const toggleApplied = () => {
    // let card know that this theme was applied,
    applyTheme(name);

    // show a lil transition for end user to show that theme was applied
    setApplied(true);
    setTimeout(() => {
      setApplied(false);
    }, 1000);
  };

  return (
    <div className={style.container}>
      <div className={style.apply}>
        {applied ? (
          <CheckIcon title="Theme applied!" className={style.applied} />
        ) : (
          <ApplyIcon
            className={style.apply}
            title="Click to apply theme!"
            onClick={toggleApplied}
          />
        )}
      </div>
      <h5 className={style.header}>{label}</h5>
      <div className={style.buttons}>
        <div className={style.minimize}></div>
        <div className={style.maximize}></div>
        <div className={style.close}></div>
      </div>
    </div>
  );
};
