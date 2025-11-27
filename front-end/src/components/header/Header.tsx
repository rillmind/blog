import { Link } from "react-router-dom";
import styles from "./Header.module.css";

function Header() {
  const headItems = [
    { path: "/login", label: "Login" },
    { path: "/posts", label: "Posts" },
  ];

  return (
    <div className={styles.header}>
      {headItems.map((item) => (
        <Link key={item.path} to={item.path} className={styles.link}>
          {item.label}
        </Link>
      ))}
    </div>
  );
}

export default Header;
