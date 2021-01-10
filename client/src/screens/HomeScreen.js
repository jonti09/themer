import { useState } from "react";
import { Pagination } from "../components/Pagination";
import { Card } from "../components/TerminalCard";
import style from "./HomeScreen.module.scss";

const itemsPerPage = 9;

export const HomeScreen = ({ themes, applyTheme }) => {
  const [activeCard, setActiveCard] = useState(-1);
  const [activePage, setActivePage] = useState(0);

  const themesForPage = themes.slice(
    activePage * itemsPerPage,
    (activePage + 1) * itemsPerPage
  );

  const card = (theme, cardIndex) => {
    return (
      <div key={theme.name} onClick={() => setActiveCard(cardIndex)}>
        <Card
          theme={theme}
          isActive={activeCard === cardIndex}
          applyTheme={applyTheme}
          className={style.card}
        />
      </div>
    );
  };

  return (
    <>
      <Pagination
        pages={themes.length / itemsPerPage}
        activePage={activePage}
        setActivePage={setActivePage}
      />
      <div className={style.container}>
        {themesForPage.map((theme, index) => card(theme, index))}
      </div>
    </>
  );
};
