import { useContext, useState } from "react";
import { AuthContext } from "../../contexts/authContext";
import Detail from "./components/Detail";
import Informasi from "./components/Informasi";
import WatchlistLayout from "../../layouts/WatchlistLayout";
import Pagination from "../../components/Pagination"
import Loading from "../../components/Loading";
import { useFetchWatchlist } from "../../hooks/useWatchlist";
import { INIT_PAGE_NUMBER } from "../../utils/constant";

const WatchListPage = () => {
    const { user } = useContext(AuthContext);
    const [page, setPage] = useState(INIT_PAGE_NUMBER)
    const { data: watchlists, loading: watchlistLoading } = useFetchWatchlist(user?.id, page);
    return (
        <div className="mx-auto max-w-4xl mt-28 space-y-6 bg-white rounded-2xl shadow-lg border border-gray-100 py-6 px-12">
            <h1 className="text-3xl font-bold text-center mb-8">🎬 Your Watchlist</h1>
            {
                watchlistLoading ? (
                    <Loading>Loading your watchlist...</Loading>
                ) : (
                    watchlists?.user_films && watchlists.user_films.map((user_film) => (
                        <WatchlistLayout key={user_film.id} watchlist={user_film}>
                            <Informasi watch={user_film} />
                            <Detail watchlist={user_film} />
                        </WatchlistLayout>
                    ))
                )
            }

            <Pagination contents={watchlists} page={page} setPage={setPage} />
        </div>
    );
};

export default WatchListPage;
