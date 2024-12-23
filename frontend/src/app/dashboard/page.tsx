"use client";
import { useEffect, useState, useRef } from "react";
import { useRouter } from "next/navigation";

interface Transaction {
  id: string;
  categoryName: string;
  amount: number;
  description: string;
  createdAt: string;
}

export default function Dashboard() {
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [categoryFilter, setCategoryFilter] = useState("All Categories");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");

  // States to control menus
  const [menuOpen1, setMenuOpen1] = useState(false);
  const [menuOpen2, setMenuOpen2] = useState(false);

  const router = useRouter();
  const menu1Ref = useRef<HTMLUListElement>(null);
  const menu2Ref = useRef<HTMLUListElement>(null);
  const button1Ref = useRef<HTMLButtonElement>(null);
  const button2Ref = useRef<HTMLButtonElement>(null);

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const response = await fetch(
          "http://localhost:8080/api/v1/users/transactions",
          {
            method: "GET",
            credentials: "include",
          },
        );
        if (!response.ok) {
          if (response.status === 401) {
            router.push("/login");
          }
          const errorData = await response.json();
          throw new Error(errorData.error || "Failed to fetch transactions");
        }
        const data = await response.json();
        setTransactions(data);
      } catch (err: unknown) {
        if (err instanceof Error) {
          setError(err.message || "Something went wrong");
        }
      }
    };
    fetchTransactions();
  }, [router]);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        menuOpen1 &&
        menu1Ref.current &&
        !menu1Ref.current.contains(event.target as Node) &&
        button1Ref.current &&
        !button1Ref.current.contains(event.target as Node)
      ) {
        setMenuOpen1(false);
      }
      if (
        menuOpen2 &&
        menu2Ref.current &&
        !menu2Ref.current.contains(event.target as Node) &&
        button2Ref.current &&
        !button2Ref.current.contains(event.target as Node)
      ) {
        setMenuOpen2(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [menuOpen1, menuOpen2]);

  // Filter logic
  const filteredTransactions = transactions.filter((transaction) => {
    const matchesCategory =
      categoryFilter === "All Categories" ||
      transaction.categoryName === categoryFilter;

    const matchesDateRange =
      (!startDate || new Date(transaction.createdAt) >= new Date(startDate)) &&
      (!endDate || new Date(transaction.createdAt) <= new Date(endDate));

    return matchesCategory && matchesDateRange;
  });

  const toggleMenu1 = () => {
    setMenuOpen1(!menuOpen1);
    if (!menuOpen1) {
      setMenuOpen2(false);
    }
  };

  const toggleMenu2 = () => {
    setMenuOpen2(!menuOpen2);
    if (!menuOpen2) {
      setMenuOpen1(false);
    }
  };

  return (
    <div className="h-screen bg-gray-900 p-6 text-gray-100">
      {/* Navbar */}
      <div className="navbar bg-gray-800 text-gray-100 shadow-lg rounded-lg mb-6 relative">
        <div className="flex-none relative">
          <button
            ref={button1Ref}
            className="btn btn-square btn-ghost"
            onClick={toggleMenu1}
          >
            {/* First menu toggle button */}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              className="inline-block h-5 w-5 stroke-current"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M4 6h16M4 12h16M4 18h16"
              ></path>
            </svg>
          </button>

          {/* Conditionally render your first menu */}
          {menuOpen1 && (
            <ul
              ref={menu1Ref}
              className="menu bg-gray-700 rounded-box w-56 absolute top-12 left-0 z-50 shadow"
            >
              <li>
                <a className="hover:bg-gray-600">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
                    />
                  </svg>
                  Transactions
                </a>
              </li>
              <li>
                <a
                  onClick={() => router.push("/categories")}
                  className="hover:bg-gray-600 cursor-pointer"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                  </svg>
                  Categories
                </a>
              </li>
              <li>
                <a className="hover:bg-gray-600">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
                    />
                  </svg>
                  Insights
                </a>
              </li>
            </ul>
          )}
        </div>
        <div className="flex-1">
          <a className="btn btn-ghost text-xl font-mono">tracker</a>
        </div>

        <div className="flex-none relative">
          <button
            ref={button2Ref}
            className="btn btn-square btn-ghost"
            onClick={toggleMenu2}
          >
            {/* Second menu toggle button */}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              className="inline-block h-5 w-5 stroke-current"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M5 12h.01M12 12h.01M19 12h.01M6 12a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0zm7 0a1 1 0 11-2 0 1 1 0 012 0z"
              ></path>
            </svg>
          </button>

          {/* Conditionally render your second menu */}
          {menuOpen2 && (
            <ul
              ref={menu2Ref}
              className="menu bg-gray-700 rounded-box w-56 absolute top-12 right-0 z-50 shadow"
            >
              <li>
                <a className="hover:bg-gray-600">Profile</a>
              </li>
              <li>
                <a className="hover:bg-gray-600">Settings</a>
              </li>
              <li>
                <a className="hover:bg-gray-600">Sign out</a>
              </li>
            </ul>
          )}
        </div>
      </div>

      <div className="max-w-4xl mx-auto bg-gray-800 rounded-lg shadow-md p-8">
        {/* Stats Section */}
        <div className="stats shadow bg-gray-800 rounded-lg p-4 mb-8 text-gray-100 flex flex-col sm:flex-row sm:justify-around">
          <div className="stat place-items-center">
            <div className="stat-title text-gray-400 font-mono">
              Total Expenses
            </div>
            <div className="stat-value text-xl font-mono font-bold">
              £{transactions.reduce((sum, t) => sum + t.amount, 0).toFixed(2)}
            </div>
            <div className="stat-desc text-gray-400 font-mono">
              Across all transactions
            </div>
          </div>
          <div className="stat place-items-center">
            <div className="stat-title text-gray-400 font-mono">
              Transactions
            </div>
            <div className="stat-value text-secondary text-xl font-mono font-bold">
              {transactions.length}
            </div>
            <div className="stat-desc text-secondary text-sm font-mono">
              Recorded entries
            </div>
          </div>
          <div className="stat place-items-center">
            <div className="stat-title text-gray-400 font-mono">
              Avg Expense
            </div>
            <div className="stat-value text-xl font-mono font-bold">
              £
              {(transactions.length > 0
                ? transactions.reduce((sum, t) => sum + t.amount, 0) /
                  transactions.length
                : 0
              ).toFixed(2)}
            </div>
            <div className="stat-desc text-gray-400 font-mono">
              Per transaction
            </div>
          </div>
        </div>
        {/* End Stats Section */}

        <div className="flex flex-col sm:flex-row justify-center items-center gap-4 mb-6">
          <input
            type="date"
            value={startDate}
            onChange={(e) => setStartDate(e.target.value)}
            className="p-2 rounded bg-gray-700 text-gray-100 font-mono focus:outline-none focus:ring-2 focus:ring-gray-500 w-full sm:w-auto"
          />
          <input
            type="date"
            value={endDate}
            onChange={(e) => setEndDate(e.target.value)}
            className="p-2 rounded bg-gray-700 text-gray-100 font-mono focus:outline-none focus:ring-2 focus:ring-gray-500 w-full sm:w-auto"
          />
          <select
            value={categoryFilter}
            onChange={(e) => setCategoryFilter(e.target.value)}
            className="p-2 rounded bg-gray-700 text-gray-100 font-mono focus:outline-none focus:ring-2 focus:ring-gray-500 w-full sm:w-auto"
          >
            <option>All Categories</option>
            <option>Bills</option>
            <option>Food</option>
            <option>Transport</option>
          </select>
        </div>

        {error && <p className="text-red-500 text-center mb-4">{error}</p>}

        {/* Add Expense Button */}
        <div className="mb-4 flex justify-end">
          <button
            onClick={() => router.push("/add-expense")}
            className="bg-gray-700 hover:bg-gray-600 text-gray-100 font-bold py-2 px-4 font-mono rounded focus:outline-none focus:shadow-outline"
          >
            Add Expense
          </button>
        </div>

        <div className="space-y-4">
          {filteredTransactions.map((transaction) => (
            <div
              key={transaction.id}
              className="flex items-center justify-between p-4 border border-gray-700 rounded-lg shadow-sm hover:shadow-lg transition-shadow bg-gray-800 hover:bg-gray-700"
            >
              <div className="flex items-center space-x-4">
                <div className="w-10 h-10 flex items-center justify-center bg-gray-700 rounded-full font-bold font-mono uppercase text-gray-100">
                  {transaction.categoryName.charAt(0)}
                </div>
                <div>
                  <p className="font-mono text-gray-100 font-bold">
                    {transaction.categoryName}
                  </p>
                  <p className="text-sm text-gray-400 font-mono">
                    {transaction.description}
                  </p>
                </div>
              </div>
              <div className="text-right">
                <p className="font-mono font-bold text-lg text-gray-100">
                  £{transaction.amount.toFixed(2)}
                </p>
                <p className="text-sm text-gray-400 font-mono">
                  {new Date(transaction.createdAt).toLocaleDateString()}
                </p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
