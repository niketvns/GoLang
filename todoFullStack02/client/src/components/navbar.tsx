import { Box, Flex, Button, Text, Container } from "@chakra-ui/react";
import { IoMoon } from "react-icons/io5";
import { LuSun } from "react-icons/lu";
import { useColorMode, useColorModeValue } from "./ui/color-mode";

export default function Navbar() {
  const { colorMode, toggleColorMode } = useColorMode();
  const bgColorValue = useColorModeValue("gray.400", "gray.700");

  return (
    <Container maxW={"900px"}>
      <Box bg={bgColorValue} px={4} my={4} borderRadius={"5"}>
        <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
          <Flex
            justifyContent={"center"}
            alignItems={"center"}
            gap={3}
            display={{ base: "none", sm: "flex" }}
          >
            <img src="/react.png" alt="logo" width={50} height={50} />
            <Text fontSize={"40px"}>+</Text>
            <img src="/go.png" alt="logo" width={40} height={40} />
            <Text fontSize={"40px"}>=</Text>
            <img src="/explode.png" alt="logo" width={50} height={50} />
          </Flex>

          <Flex alignItems={"center"} gap={3}>
            <Text fontSize={"lg"} fontWeight={500}>
              Daily Tasks
            </Text>
            <Button onClick={toggleColorMode}>
              {colorMode === "light" ? <IoMoon /> : <LuSun size={20} />}
            </Button>
          </Flex>
        </Flex>
      </Box>
    </Container>
  );
}
