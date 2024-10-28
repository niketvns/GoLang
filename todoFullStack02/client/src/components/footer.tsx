import { Link } from "@chakra-ui/react";

const Footer = () => {
  return (
    <footer
      style={{
        padding: "30px 20px",
        marginInline: "auto",
        textAlign: "center",
      }}
    >
      Made with 💖 by <Link href="/">Niket Kumar Mishra</Link>
    </footer>
  );
};

export default Footer;
