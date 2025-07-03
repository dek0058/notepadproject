
using System.Collections.Generic;

static class QuickSort<T>
{
    public static void Execute(List<T> arr, Func<T, T, bool> comparator)
    {
        if (arr.Count <= 1)
        {
            return;
        }
        Process(arr, 0, arr.Count - 1, comparator);
    }

    private static void Process(List<T> arr, int low, int high, Func<T, T, bool> comparator)
    {
        if (low >= high)
        {
            return;
        }

        var pivot = Partition(arr, low, high, comparator);
        Process(arr, low, pivot - 1, comparator);
        Process(arr, pivot + 1, high, comparator);
    }

    private static int Partition(List<T> arr, int low, int high, Func<T, T, bool> comparator)
    {
        T pivot = arr[high];
        int i = low - 1;

        for (int j = low; j < high; ++j)
        {
            if (comparator(arr[j], pivot))
            {
                ++i;
                (arr[i], arr[j]) = (arr[j], arr[i]);
            }
        }

        (arr[i + 1], arr[high]) = (arr[high], arr[i + 1]);
        return i + 1;
    }
}
