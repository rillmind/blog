import { Link } from "react-router-dom";
import styles from "./Header.module.css";

function Header() {
  const headItems = [
    // { path: "/login", label: "Login" },
    { path: "/blog", label: "Blog" },
    { path: "/journal", label: "Journal" },
    { path: "/about", label: "About" },
  ];

  return (
    <div className={styles.header}>
      <div className={styles.nav}>
        <a href="/">
          <img src="./olhar de raios2.png" alt="" className={styles.logo} />
        </a>
        <ul className={styles.navLinks}>
          {headItems.map((item) => (
            <Link key={item.path} to={item.path} className={styles.links}>
              <li className={styles.link}>{item.label}</li>
            </Link>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default Header;
