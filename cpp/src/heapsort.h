
#pragma once

#include <vector>
#include <concepts>

namespace
{
    template <typename T>
    concept Comparable = requires(T a, T b) {
        { a < b } -> std::convertible_to<bool>;
        { a > b } -> std::convertible_to<bool>;
    };

    template <Comparable T>
    void heapifyMax(std::vector<T> &arr, size_t n, size_t i)
    {
        size_t largest = i;
        size_t left = (2 * i) + 1;
        size_t right = (2 * i) + 2;

        if (left < n && arr[largest] < arr[left])
        {
            largest = left;
        }

        if (right < n && arr[largest] < arr[right])
        {
            largest = right;
        }

        if (largest != i)
        {
            T tmp = arr[largest];
            arr[largest] = arr[i];
            arr[i] = tmp;
            heapifyMax(arr, n, largest);
        }
    }

    template <Comparable T>
    void heapifyMin(std::vector<T> &arr, size_t n, size_t i)
    {
        size_t largest = i;
        size_t left = (2 * i) + 1;
        size_t right = (2 * i) + 2;

        if (left < n && arr[largest] > arr[left])
        {
            largest = left;
        }

        if (right < n && arr[largest] > arr[right])
        {
            largest = right;
        }

        if (largest != i)
        {
            T tmp = arr[largest];
            arr[largest] = arr[i];
            arr[i] = tmp;
            heapifyMin(arr, n, largest);
        }
    }
}

namespace sort
{
    template <Comparable T>
    void MaxHeapSort(std::vector<T> &arr)
    {
        size_t n = arr.size();
        for (size_t i = n / 2; i-- > 0;)
        {
            heapifyMax(arr, n, i);
        }

        for (size_t i = n - 1; i > 0; --i)
        {
            T tmp = arr[0];
            arr[0] = arr[i];
            arr[i] = tmp;

            heapifyMax(arr, i, 0);
        }
    }

    template <Comparable T>
    void MinHeapSort(std::vector<T> &arr)
    {
        size_t n = arr.size();
        for (size_t i = n / 2; i-- > 0;)
        {
            heapifyMin(arr, n, i);
        }

        for (size_t i = n - 1; i > 0; --i)
        {
            T tmp = arr[0];
            arr[0] = arr[i];
            arr[i] = tmp;

            heapifyMin(arr, i, 0);
        }
    }
}