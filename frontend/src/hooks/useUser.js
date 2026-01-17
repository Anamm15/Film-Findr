import { useMutation, useQuery } from "@tanstack/react-query";
import { getMe, getUserByUsername, logoutUser } from "../service/user";

export function useSession() {
   return useQuery({
      queryKey: ["session"],
      queryFn: async () => {
         const response = await getMe();
         return response.data;
      },
      staleTime: 1000 * 60 * 5,
      retry: false,
   })
}

export function useLogout() {
   return useMutation({
      mutationFn: logoutUser
   })
}

export function useFetchUserByUsername(username) {
   return useQuery({
      queryKey: ["user", username],
      queryFn: async () => {
         const response = await getUserByUsername(username);
         return response.data;
      },
      enabled: !!username,
      staleTime: 1000 * 60 * 5,
   });
}
