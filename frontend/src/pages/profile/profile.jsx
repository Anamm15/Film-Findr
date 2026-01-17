import { useContext, useState } from "react";
import { useParams } from "react-router-dom";
import { AuthContext } from "../../contexts/authContext";
import ReviewLayout from "../../layouts/ReviewLayout";
import Watchlist from "./components/Watchlist";
import Akun from "./components/Akun";
import { useFetchUserByUsername } from "../../hooks/useUser";
import { useFetchReviewByUserId } from "../../hooks/useReview";
import { useFetchWatchlist } from "../../hooks/useWatchlist";
import PageLoading from "../../components/PageLoading";
import { INIT_PAGE_NUMBER } from "../../utils/constant";

const ProfilePage = () => {
   const params = useParams();
   const usernameParams = params.username;
   const [reviewPage, setReviewPage] = useState(INIT_PAGE_NUMBER);
   const [watchlistPage, setWatchlistPage] = useState(INIT_PAGE_NUMBER);
   const { user: currentUser } = useContext(AuthContext);
   const finalUsername = usernameParams || currentUser?.username;

   const { data: user, loading: userLoading } = useFetchUserByUsername(finalUsername);
   const { data: reviews, loading: reviewLoading } = useFetchReviewByUserId(user?.id, reviewPage);
   const { data: watchlists, loading: watchlistLoading } = useFetchWatchlist(user?.id, watchlistPage);

   if (userLoading) {
      return <PageLoading message="Loading user" />;
   }

   return (
      <>
         <Akun user={user} review={reviews} watchlists={watchlists} icon={false} />
         <Watchlist watchlists={watchlists} watchlistPage={watchlistPage} setWatchlistPage={setWatchlistPage} loading={watchlistLoading} />
         <div className="mt-12 max-w-4xl mx-auto">
            <ReviewLayout review={reviews} setPage={setReviewPage} page={reviewPage} loading={reviewLoading} />
         </div>
      </>
   );
};

export default ProfilePage;