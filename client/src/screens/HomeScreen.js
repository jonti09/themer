import { useState } from "react";
import { Card } from "../components/terminal-card";
import style from "./HomeScreen.module.scss";

export const HomeScreen = ({ themes }) => {
  const [activeCard, setActiveCard] = useState(-1);

  console.log("active", activeCard);

  const card = (theme, cardIndex) => {
    return (
      <div key={theme.name} onClick={() => setActiveCard(cardIndex)}>
        <Card
          theme={theme}
          isActive={activeCard === cardIndex}
          className={style.card}
        />
      </div>
    );
  };

  return (
    <div className={style.container}>
      {themes.map((theme, index) => card(theme, index))}
    </div>
  );
};
