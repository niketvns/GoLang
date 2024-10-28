import "./App.css";
import { Container, Stack } from "@chakra-ui/react";
import Navbar from "./components/navbar";
import TodoForm from "./components/todoForm";
import TodoList from "./components/todoList";
import Footer from "./components/footer";

function App() {
  return (
    <Stack h="100vh">
      <Navbar />
      <Container maxWidth={"2xl"}>
        <TodoForm />
        <TodoList />
      </Container>
      <Footer />
    </Stack>
  );
}

export default App;
