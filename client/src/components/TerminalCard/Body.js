import style from "./Body.module.scss";

const topRow = [
  "black",
  "red",
  "green",
  "yellow",
  "blue",
  "purple",
  "cyan",
  "white",
];

const bottomRow = [
  "brightBlack",
  "brightRed",
  "brightGreen",
  "brightYellow",
  "brightBlue",
  "brightPurple",
  "brightCyan",
  "brightWhite",
];

const colorRow = (themeName, color, rowIndex) => {
  return (
    <div className={style.row} key={"row" + themeName + rowIndex}>
      <div className={style.block} style={{ backgroundColor: color }}></div>
      <p style={{ color: color }} className={style.text}>
        *** AaBbCc XxYyZz - Color {rowIndex + 1} ---
      </p>
      <div className={style.block} style={{ backgroundColor: color }}></div>
    </div>
  );
};

const blocks = (theme) => {
  return (
    <div className={style.blocksContainer}>
      <div className={style.blocks}>
        {topRow.map((color) => (
          <div
            key={"block" + theme.name + color}
            style={{ backgroundColor: theme[color] }}
            className={style.block}
          ></div>
        ))}
      </div>
      <div className={style.blocks}>
        {bottomRow.map((color) => (
          <div
            key={"block" + theme.name + color}
            style={{ backgroundColor: theme[color] }}
            className={style.block}
          ></div>
        ))}
      </div>
    </div>
  );
};

const applyThemeRow = (theme, isActive) => {
  return (
    <div className={style.cmdRow}>
      <span style={{ color: theme.foreground }} className={style.ini}>
        $
      </span>
      <span style={{ color: theme.red }} className={style.cmd}>
        ./apply_these.sh
      </span>
      <div
        className={style.cursor + " " + (isActive ? style.active : "")}
        style={{ backgroundColor: theme.cursorColor }}
      />
    </div>
  );
};

const colorRows = (theme, isActive) => {
  return (
    <>
      {blocks(theme)}
      {[...topRow, ...bottomRow].map((color, index) =>
        colorRow(theme["name"], theme[color], index)
      )}
      {applyThemeRow(theme, isActive)}
    </>
  );
};

export const Body = ({ theme, isActive }) => {
  return (
    <div
      style={{ backgroundColor: theme.background, color: theme.foreground }}
      className={style.container}
    >
      {colorRows(theme, isActive)}
    </div>
  );
};
