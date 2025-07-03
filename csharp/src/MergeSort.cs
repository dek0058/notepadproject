
using System.Collections.Generic;

static class MergeSort<T>
{
    public static void Execute(List<T> arr, Func<T, T, bool> comparator)
    {
        if (arr.Count <= 1)
        {
            return;
        }
        Process(arr, 0, arr.Count - 1, comparator);
    }

    private static void Process(List<T> arr, int left, int right, Func<T, T, bool> comparator)
    {
        if (left >= right)
        {
            return;
        }
        int mid = left + ((right - left) / 2);
        Process(arr, left, mid, comparator);
        Process(arr, mid + 1, right, comparator);
        Merge(arr, left, mid, right, comparator);
    }

    private static void Merge(List<T> arr, int left, int mid, int right, Func<T, T, bool> comparator)
    {
        T?[] tmp = new T?[right - left + 1];
        int i = left;
        int j = mid + 1;
        int k = 0;

        while (i <= mid && j <= right)
        {
            if (comparator(arr[i], arr[j]))
            {
                tmp[k++] = arr[i++];
            }
            else
            {
                tmp[k++] = arr[j++];
            }
        }

        for (int idx = i; idx <= mid; ++idx)
        {
            tmp[k++] = arr[idx];
        }

        for (int idx = j; idx <= right; ++idx)
        {
            tmp[k++] = arr[idx];
        }

        for (int idx = 0; idx < tmp.Length; ++idx)
        {
            var nullableItem = tmp[idx];
            if (nullableItem != null)
            {
                arr[left + idx] = nullableItem;
            }
        }
    }
}