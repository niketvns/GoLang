import { Badge, Box, Flex, GridItem, Spinner, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Todo } from "../interfaces";
import { apiBaseUrl } from "../utils/helperVariables";

const TodoItem = ({ todo }: { todo: Todo }) => {
  const queryClient = useQueryClient();

  const { mutate: updateTodo, isPending: isUpdating } = useMutation({
    mutationKey: ["updateTodo"],
    mutationFn: async () => {
      if (todo.isCompleted) return alert("Todo is already completed");
      try {
        const res = await fetch(apiBaseUrl + `/todos/${todo._id}`, {
          method: "PATCH",
        });
        const data = await res.json();
        if (!res.ok) {
          throw new Error(data.error || "Something went wrong");
        }
        return data;
      } catch (error) {
        console.log(error);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  const { mutate: deleteTodo, isPending: isDeleting } = useMutation({
    mutationKey: ["deleteTodo"],
    mutationFn: async () => {
      try {
        const res = await fetch(apiBaseUrl + `/todos/${todo._id}`, {
          method: "DELETE",
        });
        const data = await res.json();
        if (!res.ok) {
          throw new Error(data.error || "Something went wrong");
        }
        return data;
      } catch (error) {
        console.log(error);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  return (
    <GridItem colSpan={1}>
      <Flex
        flex={1}
        border={"1px"}
        borderColor={"gray.600"}
        p={4}
        borderRadius={"lg"}
        direction={"column"}
        gap={"4"}
        boxShadow={"lg"}
        justifyContent={"flex-start"}
        alignItems={"flex-start"}
        backgroundColor={"#09090B"}
      >
        {todo.isCompleted && (
          <Badge ml="1" colorPalette="green">
            Done
          </Badge>
        )}
        {!todo.isCompleted && (
          <Badge ml="1" colorPalette="yellow">
            In Progress
          </Badge>
        )}
        <Flex gap={1} direction={"column"}>
          <Text
            color={todo.isCompleted ? "green.200" : "yellow.100"}
            textDecoration={todo.isCompleted ? "line-through" : "none"}
            fontSize={"lg"}
            fontWeight={"semibold"}
          >
            {todo.title}
          </Text>
          {todo.description && (
            <Text
              color={todo.isCompleted ? "green.200" : "yellow.100"}
              textDecoration={todo.isCompleted ? "line-through" : "none"}
            >
              {todo.description}
            </Text>
          )}
        </Flex>
        <Flex gap={2} alignItems={"center"}>
          <Box
            color={"green.500"}
            cursor={"pointer"}
            onClick={() => updateTodo()}
          >
            {!isUpdating && <FaCheckCircle size={20} />}
            {isUpdating && <Spinner size={"sm"} />}
          </Box>
          <Box
            color={"red.500"}
            cursor={"pointer"}
            onClick={() => deleteTodo()}
          >
            {!isDeleting && <MdDelete size={25} />}
            {isDeleting && <Spinner size={"sm"} />}
          </Box>
        </Flex>
      </Flex>
    </GridItem>
  );
};
export default TodoItem;
