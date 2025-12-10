import { Link } from "react-router-dom";
import Header from "../../components/header/Header";
import styles from "./Blog.module.css";
import "../../Global.css";

function Blog() {
  const blogItems = [
    // Receber "itens" do back. Rota: '/blog/'
    // { path: "/login", label: "Login" },
    { path: "/about", label: "About" },
  ];

  return (
    <>
      <Header />
      <main>
        <ul className={styles.posts}>
          <h1>Blog posts</h1>
          {blogItems.map((item) => (
            <li>
              <Link key={item.path} to={item.path} className={styles.links}>
                <p>{item.label}</p>
              </Link>
              <div className={styles.postDate}>
                <time>Dec 01, 2025</time>
              </div>
            </li>
          ))}
        </ul>
      </main>
    </>
  );
}

export default Blog;
