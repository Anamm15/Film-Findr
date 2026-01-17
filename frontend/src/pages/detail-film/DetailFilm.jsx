import { useParams } from "react-router-dom";
import WatchListForm from "./components/AddWatchlist";
import RekomendasiFilm from "./components/Rekomendasi";
import AddReview from "./components/AddReview";
import Sinopsis from "./components/Sinopsis";
import ReviewLayout from "../../layouts/ReviewLayout";
import InformasiFilm from "./components/InformasiFilm";
import Gambar from "./components/Gambar";
import { useFetchFilmById } from "../../hooks/useFilm";
import { useReviewsByFilm } from "../../hooks/useReview";
import PageLoading from "../../components/PageLoading";
import { useState, useEffect } from "react";
import { INIT_PAGE_NUMBER } from "../../utils/constant";

const DetailFilmPage = () => {
   const { id } = useParams();
   const [page, setPage] = useState(INIT_PAGE_NUMBER);
   const { data: films, loading: loadingFilms } = useFetchFilmById(id);
   const { data, loading: loadingReviews } = useReviewsByFilm(id, page);
   const [reviews, setReviews] = useState();

   useEffect(() => {
      if (!data) return
      setReviews(data)
   }, [data, reviews]);

   if (loadingFilms) {
      return <PageLoading message="Loading film..." />;
   }

   return (
      <div className="max-w-5xl mx-auto px-4 pb-10 bg-background">
         {
            films && (
               <>
                  <div className="flex flex-col md:flex-row gap-6 my-8 shadow-md rounded-xl p-4 relative mt-28">
                     <Gambar film={films} />
                     <div>
                        <InformasiFilm film={films} />
                        <WatchListForm id={id} />
                     </div>
                  </div>

                  <Sinopsis sinopsis={films.sinopsis} />
                  <ReviewLayout review={reviews} setReviews={setReviews} setPage={setPage} page={page} loading={loadingReviews} />
                  <AddReview id={id} setReviews={setReviews} />
                  {
                     films && <RekomendasiFilm films={films.films} />
                  }
               </>
            )
         }
      </div>
   );
};

export default DetailFilmPage;