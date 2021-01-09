import style from "./Pagination.module.scss";

export const Pagination = ({ pages, activePage, setActivePage }) => {
  pages = [...Array(pages).keys()];

  return (
    <div className={style.container}>
      <div
        onClick={() => setActivePage(activePage > 0 ? activePage - 1 : 0)}
        className={style.arrow}
      >
        {"<"}
      </div>
      {pages.map((page) => (
        <div
          key={page}
          className={page === activePage ? style.active : ""}
          onClick={() => setActivePage(page)}
        >
          {page + 1}
        </div>
      ))}
      <div
        onClick={() =>
          setActivePage(activePage < pages.length ? activePage + 1 : 0)
        }
        className={style.arrow}
      >
        {">"}
      </div>
    </div>
  );
};
