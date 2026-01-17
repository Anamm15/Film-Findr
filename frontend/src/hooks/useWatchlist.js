import { useQuery } from "@tanstack/react-query";
import { INIT_PAGE_NUMBER } from "../utils/constant";
import { getUserFilmByUserId } from "../service/userFilm";

export function useFetchWatchlist(userId, page = INIT_PAGE_NUMBER) {
   return useQuery({
      queryKey: ["watchlists", userId, page],
      queryFn: async () => {
         const response = await getUserFilmByUserId(userId, page);
         return response.data;
      },
      enabled: !!userId,
      staleTime: 1000 * 60 * 5,
      keepPreviousData: true,
   });
}
