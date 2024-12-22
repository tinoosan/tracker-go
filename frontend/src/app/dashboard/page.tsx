"use client";

import { useEffect, useState } from "react";
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
  const router = useRouter();

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

  return (
    <div className="h-screen bg-gray-900 p-6 overflow-y-auto">
      <div className="max-w-4xl mx-auto bg-gray-800 rounded-lg shadow-md p-8">
        <h1 className="text-gray-100 text-2xl font-bold font-mono text-center mb-6">
          tracker
        </h1>

        {error && <p className="text-red-500 text-center mb-4">{error}</p>}
        <div className="space-y-4">
          {transactions.map((transaction) => (
            <div
              key={transaction.id}
              className="flex items-center justify-between p-4 border border-gray-700 rounded-lg shadow-sm hover:shadow-md bg-gray-800"
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
