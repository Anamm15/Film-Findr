import WatchlistLayout from "../../../layouts/WatchlistLayout";
import Informasi from "./Informasi";
import Pagination from "../../../components/Pagination";
import Loading from "../../../components/Loading";

const Watchlist = (props) => {
    const { watchlists, watchlistPage, setWatchlistPage, loading } = props;
    return (
        <div className="mt-12 px-4 max-w-4xl mx-auto space-y-4 bg-white rounded-2xl shadow-lg p-4 hover:shadow-xl transition-shadow">
            <h2 className="text-3xl font-semibold ps-4 pt-4">Watchlist</h2>
            {
                loading ? (
                    <Loading>Loading your watchlist...</Loading>
                ) : (
                    watchlists?.user_films && watchlists.user_films.map((user_film) => (
                        <WatchlistLayout key={user_film.id} watchlist={user_film}>
                            <Informasi watch={user_film} />
                        </WatchlistLayout>
                    ))
                )
            }

            <Pagination contents={watchlists} page={watchlistPage} setPage={setWatchlistPage} />
        </div>
    )
}

export default Watchlist;