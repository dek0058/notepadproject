
using System.Collections.Generic;

static class MaxHeapSort<T> where T : IComparable<T>
{
    static public void Execute(List<T> arr)
    {
        for (int idx = (arr.Count / 2) - 1; idx >= 0; --idx)
        {
            HeapSort(arr, arr.Count, idx);
        }

        for (int idx = arr.Count - 1; idx > 0; --idx)
        {
            (arr[0], arr[idx]) = (arr[idx], arr[0]);
            HeapSort(arr, idx, 0);
        }
    }

    static private void HeapSort(List<T> arr, int cnt, int idx)
    {
        int largest = idx;
        int left = (2 * idx) + 1;
        int right = (2 * idx) + 2;

        if (left < cnt && arr[largest].CompareTo(arr[left]) < 0)
        {
            largest = left;
        }

        if (right < cnt && arr[largest].CompareTo(arr[right]) < 0)
        {
            largest = right;
        }

        if (largest != idx)
        {
            (arr[largest], arr[idx]) = (arr[idx], arr[largest]);
            HeapSort(arr, cnt, largest);
        }
    }
}

static class MinHeapSort<T> where T : IComparable<T>
{
    static public void Execute(List<T> arr)
    {
        for (int idx = (arr.Count / 2) - 1; idx >= 0; --idx)
        {
            HeapSort(arr, arr.Count, idx);
        }

        for (int idx = arr.Count - 1; idx > 0; --idx)
        {
            (arr[0], arr[idx]) = (arr[idx], arr[0]);
            HeapSort(arr, idx, 0);
        }
    }

    static private void HeapSort(List<T> arr, int cnt, int idx)
    {
        int largest = idx;
        int left = (2 * idx) + 1;
        int right = (2 * idx) + 2;

        if (left < cnt && arr[largest].CompareTo(arr[left]) > 0)
        {
            largest = left;
        }

        if (right < cnt && arr[largest].CompareTo(arr[right]) > 0)
        {
            largest = right;
        }

        if (largest != idx)
        {
            (arr[largest], arr[idx]) = (arr[idx], arr[largest]);
            HeapSort(arr, cnt, largest);
        }
    }
}