export function SkeletonCard() {
  return (
    <div className="bg-white rounded-2xl border border-gray-100 shadow-sm overflow-hidden animate-pulse">
      <div className="p-6 flex flex-col gap-3">
        <div className="flex items-center justify-between">
          <div className="h-5 w-20 bg-gray-100 rounded-full" />
          <div className="h-4 w-6 bg-gray-100 rounded" />
        </div>
        <div className="space-y-2">
          <div className="h-4 bg-gray-100 rounded w-full" />
          <div className="h-4 bg-gray-100 rounded w-4/5" />
        </div>
        <div className="space-y-1.5">
          <div className="h-3 bg-gray-100 rounded w-full" />
          <div className="h-3 bg-gray-100 rounded w-3/4" />
        </div>
        <div className="flex gap-1.5">
          <div className="h-5 w-14 bg-gray-100 rounded-full" />
          <div className="h-5 w-16 bg-gray-100 rounded-full" />
          <div className="h-5 w-12 bg-gray-100 rounded-full" />
        </div>
      </div>
      <div className="px-6 py-4 border-t border-gray-50 bg-gray-50/50 flex items-center justify-between">
        <div className="h-3 w-32 bg-gray-100 rounded" />
        <div className="h-3 w-12 bg-gray-100 rounded" />
      </div>
    </div>
  );
}
