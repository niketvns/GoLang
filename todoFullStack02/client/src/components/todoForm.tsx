/* eslint-disable @typescript-eslint/no-explicit-any */
import {
  Button,
  Flex,
  Heading,
  Input,
  Spinner,
  Textarea,
} from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { MutableRefObject, useEffect, useRef, useState } from "react";
import { IoMdAdd } from "react-icons/io";
import { apiBaseUrl } from "../utils/helperVariables";

const TodoForm = () => {
  const [newTodo, setNewTodo] = useState({
    title: "",
    description: "",
  });
  const inputRef: MutableRefObject<HTMLInputElement | null> = useRef(null);

  const queryClient = useQueryClient();

  const { mutate: createTodo, isPending: isCreating } = useMutation({
    mutationKey: ["createTodo"],
    mutationFn: async (e: React.FormEvent) => {
      e.preventDefault();
      try {
        const res = await fetch(apiBaseUrl + `/todos`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(newTodo),
        });
        const data = await res.json();

        if (!res.ok) {
          throw new Error(data.error || "Something went wrong");
        }

        setNewTodo({
          title: "",
          description: "",
        });
        return data;
      } catch (error: any) {
        throw new Error(error);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error: any) => {
      alert(error.message);
    },
  });

  useEffect(() => {
    if (inputRef.current) {
      inputRef.current?.focus();
    }
  }, []);

  return (
    <form onSubmit={createTodo}>
      <Flex
        gap={2}
        flexDirection={"column"}
        alignItems={"flex-start"}
        boxShadow={"2xl"}
        p={5}
        borderRadius={5}
      >
        <Heading>Create Todo</Heading>
        <Input
          type="text"
          value={newTodo.title}
          onChange={(e) =>
            setNewTodo((prev) => ({ ...prev, title: e.target.value }))
          }
          required
          placeholder="Enter todo title"
          ref={inputRef}
        />
        <Textarea
          rows={4}
          placeholder="Enter todo description..."
          value={newTodo.description}
          onChange={(e) =>
            setNewTodo((prev) => ({ ...prev, description: e.target.value }))
          }
        />
        <Button
          mx={2}
          type="submit"
          _active={{
            transform: "scale(.97)",
          }}
          alignSelf={"flex-end"}
        >
          Add {isCreating ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
        </Button>
      </Flex>
    </form>
  );
};
export default TodoForm;
