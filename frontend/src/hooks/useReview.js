import { useQuery } from "@tanstack/react-query";
import { getReviewByFilmId, getReviewByUserId } from "../service/review";
import { INIT_PAGE_NUMBER } from "../utils/constant";

export function useReviewsByFilm(id, page) {
   return useQuery({
      queryKey: ["reviews", id, page],
      queryFn: () => getReviewByFilmId(id, page),
      enabled: !!id,
      keepPreviousData: true,
      staleTime: 2 * 60 * 1000,
   });
}

export function useFetchReviewByUserId(id, page = INIT_PAGE_NUMBER) {
   return useQuery({
      queryKey: ["reviews", "user", id, page],
      queryFn: async () => {
         const response = await getReviewByUserId(id, page);
         return response.data;
      },
      enabled: !!id,
      staleTime: 1000 * 60 * 5,
      keepPreviousData: true,
   });
}
