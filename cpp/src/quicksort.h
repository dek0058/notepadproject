
#pragma once

#include <vector>
#include <functional>

namespace
{
    template <typename T>
    int partition(std::vector<T> &arr, int low, int high, std::function<bool(const T &, const T &)> comparator)
    {
        T pivot = arr[high];
        int i = low - 1;

        for (int j = low; j < high; j++)
        {
            if (comparator(arr[j], pivot))
            {
                ++i;
                T temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }
        }

        T temp = arr[i + 1];
        arr[i + 1] = arr[high];
        arr[high] = temp;
        return i + 1;
    }

    template <typename T>
    void quickSort(std::vector<T> &arr, int low, int high, std::function<bool(const T &, const T &)> comparator)
    {
        if (low >= high)
        {
            return;
        }

        int pivot = partition(arr, low, high, comparator);
        quickSort(arr, low, pivot - 1, comparator);
        quickSort(arr, pivot + 1, high, comparator);
    }
}

namespace sort
{
    template <typename T>
    void QuickSort(std::vector<T> &arr, std::function<bool(const T &, const T &)> comparator = std::less<T>())
    {
        if (arr.size() <= 1)
        {
            return;
        }

        quickSort(arr, 0, arr.size() - 1, comparator);
    }
}
