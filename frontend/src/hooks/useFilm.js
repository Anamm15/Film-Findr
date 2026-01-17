import { useQuery } from "@tanstack/react-query";
import { getAllFilm, getFilmById, getTopFilm, getTrendingFilm } from "../service/film";
import { INIT_PAGE_NUMBER } from "../utils/constant";


export function useFetchFilms(page = INIT_PAGE_NUMBER) {
   return useQuery({
      queryKey: ["films", page],
      queryFn: async () => {
         const response = await getAllFilm(page);
         return response.data;
      },
      staleTime: 1000 * 60 * 5,
      keepPreviousData: true,
   });
}

export function useFetchFilmById(id) {
   return useQuery({
      queryKey: ["film", id],
      queryFn: async () => {
         const response = await getFilmById(id);
         return response.data;
      },
      enabled: !!id,
      staleTime: 1000 * 60 * 5,
   });
}

export function useFetchTopFilms() {
   return useQuery({
      queryKey: ["top-films"],
      queryFn: async () => {
         const response = await getTopFilm();
         return response.data;
      },
      staleTime: 1000 * 60 * 5,
   });
}

export function useFetchTrendingFilms() {
   return useQuery({
      queryKey: ["trending-films"],
      queryFn: async () => {
         const response = await getTrendingFilm();
         return response.data;
      },
      staleTime: 1000 * 60 * 5,
   });
}

